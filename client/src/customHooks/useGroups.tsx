import { useCallback, useMemo } from "react";
import { setCurrentGroup, setTotalPages } from "../store/group.slice";
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

  const hasMoreGroups = useMemo(() => {
    return totalPages == null || currentPage + 1 <= totalPages;
  }, [currentPage, totalPages]);

  const fetchAndSetCurrentGroup = async (groupId: string) => {
    const groupDetails = await apiClientObj.get(
      routes.GET_GROUP_DETAILS.replace("{0}", groupId),
    );
    dispatch(setCurrentGroup(groupDetails));
  };

  return {
    groups,
    fetchGroupsNext,
    hasMoreGroups,
    fetchAndSetCurrentGroup,
  };
};

export default useGroups;
