import DeleteModal from "../Modal/DeleteModal/DeleteModal";
import PreviewModal from "../Modal/PreviewModal/PreviewModal";
import type { ImageInterface } from "../../common/interfaces";
import AddImageToGroup from "../Modal/AddImageToGroupModal/AddImageToGroup";
import GroupViewButtons from "./GroupViewButtons/GroupViewButtons";
import ImageViewButtons from "./ImageViewButtons/ImageViewButtons";

interface CardProps {
  imgData: ImageInterface;
  isGroupView: boolean;
}

const Card = ({ imgData, isGroupView }: CardProps) => {
  return (
    <>
      <DeleteModal />
      <PreviewModal />
      <AddImageToGroup />
      <div className="group break-inside-avoid mb-4 relative">
        <img
          className="rounded-sm w-full hover:shadow-xl min-h-[200px] h-auto block"
          src={imgData.url}
        />
        <div className="absolute bottom-0 right-0 hidden group-hover:flex w-full bg-black/60 justify-end">
          {isGroupView ? (
            <GroupViewButtons imgData={imgData} />
          ) : (
            <ImageViewButtons imgData={imgData} />
          )}
        </div>
      </div>
    </>
  );
};

export default Card;
