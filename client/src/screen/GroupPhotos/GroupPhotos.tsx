import { useDispatch } from "react-redux";
import { useNavigate, useParams } from "react-router";
import { useEffect } from "react";
import ImageList from "../../component/ImageList/ImageList";
import { setCurrentGroup } from "../../store/group.slice";
import useGroups from "../../customHooks/useGroups";
const GroupPhotos = () => {
  const params = useParams();
  const navigate = useNavigate();
  const dispatch = useDispatch();

  const { groupId } = params;

  const { groups, currentGroup, fetchGroupsPrev, fetchGroupImagesPrev } =
    useGroups();

  useEffect(() => {
    const getGroupData = async () => {
      if (!groups) {
        await fetchGroupsPrev();
      }
      if (groups && currentGroup?.id !== groupId) {
        dispatch(setCurrentGroup(groupId!));
      }
      if (
        !currentGroup?.imageData.images ||
        currentGroup?.imageData.images.length === 0
      ) {
        await fetchGroupImagesPrev();
      }
    };
    getGroupData();
  }, [
    currentGroup,
    dispatch,
    fetchGroupImagesPrev,
    fetchGroupsPrev,
    groupId,
    groups,
  ]);

  return (
    <>
      {currentGroup?.imageData && (
        <div className="flex-1 justify-center w-full">
          <div className="m-4 p-8 d-flex rounded-2xl content-center text-center border-2 border-dashed border-gray-400 bg-blue-100">
            <h1 className="text-4xl font-bold mb-4">{currentGroup?.name}</h1>

            {currentGroup?.imageData.images &&
              currentGroup?.imageData.images.length == 0 && (
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

          <ImageList
            images={currentGroup?.imageData.images || []}
            isGroupView={true}
          />
        </div>
      )}
    </>
  );
};

export default GroupPhotos;
