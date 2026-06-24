import { useCallback } from "react";
import { setTotalPages } from "../store/group.slice";
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
  const { groups, totalPages, currentPage, currentLimit } = useSelector(
    (state: RootState) => state.group,
  );
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

  const fetchGroupsNext = useCallback(async () => {
    const nextPage = currentPage + 1;
    dispatch(setCurrentPage(nextPage));
    const response = await fetchGroups(nextPage, currentLimit);
    dispatch(setGroups([...(groups || []), ...response.data]));
  }, [groups, fetchGroups, currentPage, currentLimit, dispatch]);

  const hasMoreGroups = () => {
    console.log(
      `Here are totalpage: ${totalPages} currentPage + 1: ${currentPage + 1}`,
    );
    return totalPages == null || currentPage + 1 <= totalPages;
  };

  return {
    groups,
    fetchGroupsNext,
    hasMoreGroups,
  };
};

export default useGroups;
