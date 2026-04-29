import { useDispatch, useSelector } from "react-redux";
import type { RootState } from "../../store/store";
import { useNavigate, useParams } from "react-router";
import { useCallback, useEffect } from "react";
import ImageList from "../../component/ImageList/ImageList";
import {
  setCurrentGroup,
  setFetchGroups,
  setGroupImageData,
} from "../../store/group.slice";
import { apiClientObj } from "../../common/apiClient";
import { routes } from "../../common/routes";
const GroupPhotos = () => {
  const params = useParams();
  const navigate = useNavigate();
  const dispatch = useDispatch();

  const { groupId } = params;

  const group = useSelector((state: RootState) => state.group);

  useEffect(() => {
    if (group.groups == null) {
      dispatch(setFetchGroups());
    }
  }, [dispatch, group.groups]);

  useEffect(() => {
    if (group.groups && groupId && group.currentGroup?.id !== groupId) {
      dispatch(setCurrentGroup({ id: groupId }));
    }
  }, [groupId, group, dispatch]);

  const imageData = group.currentGroup?.imageData;

  const fetchGroupImageData = useCallback(
    async (groupId: string) => {
      const currentGroup = group.currentGroup;
      if (
        currentGroup &&
        (groupId !== currentGroup.id ||
          (groupId == currentGroup.id &&
            (currentGroup.imageData.images?.length ?? 0) == 0))
      ) {
        const res = await apiClientObj.get(
          routes.GET_GROUP_IMAGES.replace("{0}", groupId),
        );

        dispatch(setGroupImageData(res));
      }
    },
    [dispatch, group],
  );

  useEffect(() => {
    if (
      group.currentGroup &&
      group.currentGroup.id &&
      (group.currentGroup.imageData.fetchImages ?? false) === true
    ) {
      fetchGroupImageData(group.currentGroup.id);
    }
  }, [group, fetchGroupImageData]);

  return (
    <>
      {imageData && (
        <div className="flex-1 justify-center w-full">
          <div className="m-4 p-8 d-flex rounded-2xl content-center text-center border-2 border-dashed border-gray-400 bg-blue-100">
            <h1 className="text-4xl font-bold mb-4">
              {group.currentGroup?.name}
            </h1>

            {imageData.images && imageData.images.length == 0 && (
              <h2 className="mb-2">No photos yet...</h2>
            )}
            <h2> Add some photos to group form photos tab</h2>
            <button
              onClick={() => navigate("/dashboard/photos")}
              className="mt-4 px-4 py-2 bg-blue-500 text-white rounded"
            >
              Goto photos Tab
            </button>
          </div>

          <ImageList imageData={imageData} isGroupView={true} />
        </div>
      )}
    </>
  );
};

export default GroupPhotos;
