import { routes } from "../../common/routes";

interface CardProps {
  imgData: {
    key: string;
  };
}

const Card = ({ imgData }: CardProps) => {
  return (
    <div className="break-inside-avoid mb-4">
      <img
        className="rounded-sm w-full hover:shadow-xl  h-auto block "
        src={routes.IMAGE_PREFIX + imgData.key}
      />
    </div>
  );
};

export default Card;
