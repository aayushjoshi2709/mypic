import { useCallback } from "react";
import { setTotalPages } from "../store/image.slice";
import {
  setCurrentPage,
  setCurrentLimit,
  setGroups,
  setCurrentGroupImages,
  setCurrentGroupImagesTotalPages,
  setCurrentGroupImagesCurrentPage,
  setCurrentGroupImagesCurrentLimit,
} from "../store/group.slice";
import { useDispatch, useSelector } from "react-redux";
import type { RootState } from "../store/store";
import { apiClientObj } from "../common/apiClient";
import { routes } from "../common/routes";

const useGroups = () => {
  const { groups, totalPages, currentGroup, currentPage, currentLimit } =
    useSelector((state: RootState) => state.group);
  const dispatch = useDispatch();

  const fetchGroups = useCallback(
    async (page: number, limit: number) => {
      dispatch(setCurrentPage(page));
      dispatch(setCurrentLimit(limit));
      const response = await apiClientObj.get(
        routes.GET_ALL_GROUPS + `?page=${page}&limit=${limit}`,
      );
      dispatch(setTotalPages(response.totalPages));
      return response;
    },
    [dispatch],
  );

  const fetchGroupNext = useCallback(async () => {
    const nextPage =
      currentPage + 1 <= totalPages! ? currentPage + 1 : totalPages!;
    dispatch(setCurrentPage(nextPage));
    const response = await fetchGroups(nextPage, currentLimit);
    dispatch(setGroups([...(groups || []), ...response.data]));
  }, [groups, fetchGroups, currentPage, currentLimit, totalPages, dispatch]);

  const fetchGroupPrev = useCallback(async () => {
    const prevPage = currentPage - 1 > 0 ? currentPage - 1 : 0;
    dispatch(setCurrentPage(prevPage));
    const response = await fetchGroups(prevPage, currentLimit);
    dispatch(setGroups([...response.data, ...(groups || [])]));
  }, [groups, fetchGroups, currentPage, currentLimit, dispatch]);

  const fetchGroupImages = useCallback(
    async (page: number = 1, limit: number = 10) => {
      const url = `${routes.GET_ALL_IMAGES}?page=${page}&limit=${limit}&groupId=${currentGroup?.id}`;
      dispatch(setCurrentGroupImagesCurrentPage(page));
      dispatch(setCurrentGroupImagesCurrentLimit(limit));
      const response = await apiClientObj.get(url);
      dispatch(setCurrentGroupImagesTotalPages(response.totalPages));
      return response;
    },
    [dispatch, currentGroup?.id],
  );

  const fetchGroupImagesNext = useCallback(async () => {
    const groupImagesPage = currentGroup?.imageData?.currentPage ?? 1;
    const nextPage = groupImagesPage + 1;
    if (nextPage > 0) {
      const response = await fetchGroupImages(nextPage, currentLimit);
      if (currentGroup?.imageData) {
        dispatch(
          setCurrentGroupImages({
            ...currentGroup?.imageData,
            images: [
              ...(currentGroup?.imageData.images || []),
              ...response.data,
            ],
          }),
        );
      }
    }
  }, [fetchGroupImages, currentLimit, currentGroup?.imageData, dispatch]);

  const fetchGroupImagesPrev = useCallback(async () => {
    const groupImagesPage = currentGroup?.imageData?.currentPage ?? 1;
    const prevPage = groupImagesPage - 1;
    if (prevPage > 0) {
      const response = await fetchGroupImages(prevPage, currentLimit);
      if (currentGroup?.imageData) {
        dispatch(
          setCurrentGroupImages({
            ...currentGroup?.imageData,
            images: [
              ...response.data,
              ...(currentGroup?.imageData.images || []),
            ],
          }),
        );
      }
    }
  }, [fetchGroupImages, currentLimit, currentGroup?.imageData, dispatch]);

  return {
    groups,
    currentGroup,
    fetchGroupImagesNext,
    fetchGroupImagesPrev,
    fetchGroupNext,
    fetchGroupPrev,
  };
};

export default useGroups;
