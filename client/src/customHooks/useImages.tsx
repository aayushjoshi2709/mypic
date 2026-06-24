import { useCallback } from "react";
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
  const { images, currentPage, currentLimit, totalPages } = useSelector(
    (state: RootState) => state.image,
  );
  const { currentGroup } = useSelector((state: RootState) => state.group);
  const dispatch = useDispatch();
  const fetchImages = useCallback(
    async (page: number = 1, limit: number = 10) => {
      let url = `${routes.GET_ALL_IMAGES}?page=${page}&limit=${limit}`;
      const groupId = currentGroup?.id;
      if (groupId) {
        url = `${url}&groupId=${groupId}`;
      }
      const response = await apiClientObj.get(url);
      dispatch(setTotalPages(response.totalPages));
      dispatch(setCurrentPage(response.page));
      dispatch(setCurrentLimit(response.limit));
      dispatch(appendImages({ images: response.data }));
    },
    [dispatch, currentGroup],
  );

  const fetchNextPage = useCallback(() => {
    fetchImages(currentPage + 1, currentLimit);
  }, [fetchImages, currentPage, currentLimit]);

  const hasMoreImages = () => {
    return totalPages == null || currentPage + 1 <= totalPages;
  };

  return {
    fetchImages,
    fetchNextPage,
    hasMoreImages,
    images,
  };
};

export default useImages;
