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
    async (url: string) => {
      const response = await apiClientObj.get(url);
      dispatch(setTotalPages(response.totalPages));
      dispatch(setCurrentPage(response.page));
      dispatch(setCurrentLimit(response.limit));
      dispatch(appendImages(response.data));
    },
    [dispatch],
  );

  const fetchNextPage = useCallback(() => {
    const url = `${routes.GET_ALL_IMAGES}?page=${currentPage + 1}&limit=${currentLimit}`;
    fetchImages(url);
  }, [fetchImages, currentPage, currentLimit]);

  const fetchNextGroupPage = useCallback(() => {
    const groupId = currentGroup?.id ?? "";
    const url = `${routes.GET_GROUP_IMAGES.replace("{0}", groupId)}?page=${currentPage + 1}&limit=${currentLimit}`;
    fetchImages(url);
  }, [currentGroup?.id, currentPage, currentLimit, fetchImages]);

  const hasMoreImages = () => {
    return totalPages == null || currentPage + 1 <= totalPages;
  };

  return {
    fetchImages,
    fetchNextPage,
    hasMoreImages,
    fetchNextGroupPage,
    images,
  };
};

export default useImages;
