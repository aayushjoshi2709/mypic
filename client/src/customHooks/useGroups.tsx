import { useCallback } from "react";
import { appendImages, setTotalPages } from "../store/image.slice";
import {
  setCurrentPage,
  setCurrentLimit,
  setGroups,
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
    async (currentPage: number, currentLimit: number) => {
      setCurrentPage(currentPage);
      setCurrentLimit(currentLimit);
      return apiClientObj.get(
        routes.GET_ALL_GROUPS + `?page=${currentPage}&limit=${currentLimit}`,
      );
    },
    [],
  );

  const fetchGroupNext = useCallback(async () => {
    if (currentPage + 1 <= totalPages!) {
      setCurrentPage(currentPage + 1);
    } else {
      setCurrentPage(totalPages!);
    }
    const response = await fetchGroups(currentPage + 1, currentLimit);
    setGroups([...(groups || []), ...response.data]);
  }, [groups, fetchGroups, currentPage, currentLimit, totalPages]);

  const fetchGroupPrev = useCallback(async () => {
    if (currentPage - 1 > 0) {
      setCurrentPage(currentPage - 1);
    } else {
      setCurrentPage(0);
    }
    const response = await fetchGroups(currentPage - 1, currentLimit);
    setGroups([...response.data, ...(groups || [])]);
  }, [groups, fetchGroups, currentPage, currentLimit]);

  const fetchGroupImages = useCallback(
    async (
      page: number = 1,
      limit: number = 10,
      arrange: "before" | "after",
    ) => {
      if (page < 0 || (totalPages && page > totalPages)) {
        throw Error("Page does not exist");
      }

      let url = `${routes.GET_ALL_IMAGES}?page=${page}&limit=${limit}`;
      if (currentGroup?.id) {
        url = `${url}&groupId=${currentGroup.id}`;
      }
      const response = await apiClientObj.get(url);
      dispatch(setTotalPages(response.totalPages));
      dispatch(setCurrentPage(response.page));
      dispatch(setCurrentLimit(response.limit));
      dispatch(appendImages({ images: response.data, arrange: arrange }));
    },
    [totalPages, dispatch, currentGroup?.id],
  );

  const fetchGroupImagesNext = useCallback(async () => {
    if (currentPage + 1 > 0) {
      const response = await fetchGroupImages(
        currentPage + 1,
        currentLimit,
        "after",
      );
    }
  }, [fetchGroupImages, currentPage, currentLimit]);

  const fetchGroupImagesPrev = useCallback(async () => {
    if (currentPage - 1 > 0) {
      const response = await fetchGroupImages(
        currentPage - 1,
        currentLimit,
        "before",
      );
    }
  }, [fetchGroupImages, currentPage, currentLimit]);

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
