import { useDispatch, useSelector } from "react-redux";
import { useNavigate, useParams } from "react-router";
import { useEffect, useState } from "react";
import ImageList from "../../component/ImageList/ImageList";
import InfiniteScroll from "react-infinite-scroll-component";
import useImages from "../../customHooks/useImages";
import type { RootState } from "../../store/store";
import {
  clearImages,
  setCurrentPage,
  setTotalPages,
} from "../../store/image.slice";
import useGroups from "../../customHooks/useGroups";
const GroupPhotos = () => {
  const params = useParams();
  const navigate = useNavigate();
  const dispatch = useDispatch();

  const { currentGroup } = useSelector((state: RootState) => state.group);
  const { fetchNextGroupPage, images, hasMoreImages } = useImages();
  const { fetchAndSetCurrentGroup } = useGroups();
  const [isLoading, setIsLoading] = useState(false);
  useEffect(() => {
    const { groupId } = params;
    const fetchGroupDetails = async (groupId: string) => {
      await fetchAndSetCurrentGroup(groupId);
      dispatch(clearImages());
      dispatch(setCurrentPage(0));
      dispatch(setTotalPages(null));
    };
    if (groupId) {
      setIsLoading(true);
      fetchGroupDetails(groupId);
      setIsLoading(false);
    }
  }, [dispatch, fetchAndSetCurrentGroup, params]);

  return (
    <>
      {currentGroup && (
        <div className="flex-1 justify-center w-full">
          <div className="m-4 p-8 d-flex rounded-2xl content-center text-center border-2 border-dashed border-gray-400 bg-blue-100">
            <div className="flex justify-center mb-8">
              <img
                src={currentGroup?.imageUrl}
                alt={currentGroup?.name}
                className="w-50 h-50 object-cover rounded-full"
              />
            </div>
            <h1 className="text-4xl rou font-bold mb-4">{currentGroup?.name}</h1>

            {images?.length == 0 && <h2 className="mb-2">No photos yet...</h2>}
            <h2> Add some photos to group form photos tab</h2>
            <button
              onClick={() => navigate("/dashboard/photos")}
              className="mt-4 px-4 py-2 bg-blue-500 text-white rounded"
            >
              Goto photos Tab
            </button>
          </div>

          <InfiniteScroll
            dataLength={images?.length ?? 0}
            next={fetchNextGroupPage}
            hasMore={!isLoading && hasMoreImages}
            loader={<p>Loading...</p>}
          >
            <ImageList images={images ?? []} isGroupView={true} />
          </InfiniteScroll>
        </div>
      )}
    </>
  );
};

export default GroupPhotos;
