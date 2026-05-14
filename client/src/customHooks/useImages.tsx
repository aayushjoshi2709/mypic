import { useCallback, useState } from "react";
import {
  appendImages,
  setCurrentLimit,
  setCurrentPage,
  setTotalPages,
} from "../store/image.slice";
import { useDispatch, useSelector } from "react-redux";
import type { RootState } from "../store/store";
import { apiClientObj } from "../common/apiClient";
import { routes } from "../common/routes";

const useImages = () => {
  const [groupId, setGroupId] = useState<string>("");
  const { images, currentPage, currentLimit, totalPages } = useSelector(
    (state: RootState) => state.image,
  );
  const dispatch = useDispatch();
  const fetchImages = useCallback(
    async (
      page: number = 1,
      limit: number = 10,
      arrange: "before" | "after",
    ) => {
      if (page < 0 || (totalPages && page > totalPages)) {
        throw Error("Page does not exist");
      }

      let url = `${routes.GET_ALL_IMAGES}?page=${page}&limit=${limit}`;
      if (groupId) {
        url = `${url}&groupId=${groupId}`;
      }
      const response = await apiClientObj.get(url);
      dispatch(setTotalPages(response.totalPages));
      dispatch(setCurrentPage(response.page));
      dispatch(setCurrentLimit(response.limit));
      dispatch(appendImages({ images: response.data, arrange: arrange }));
    },
    [totalPages, dispatch, groupId],
  );

  const fetchNextPage = () => {
    if (currentPage + 1 > 0) {
      fetchImages(currentPage + 1, currentLimit, "after");
    }
  };
  const fetchPrevPage = () => {
    if (currentPage - 1 > 0) {
      fetchImages(currentPage - 1, currentLimit, "before");
    }
  };

  return {
    groupId,
    setGroupId,
    fetchImages,
    fetchNextPage,
    fetchPrevPage,
    images,
  };
};

export default useImages;
