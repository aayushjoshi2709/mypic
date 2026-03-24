interface CardProps {
  imgData: {
    src: string;
  };
}

const Card = ({ imgData }: CardProps) => {
  return (
    <div className="break-inside-avoid mb-4">
      <img className="rounded-sm w-full  h-auto block" src={imgData.src} />
    </div>
  );
};

export default Card;
