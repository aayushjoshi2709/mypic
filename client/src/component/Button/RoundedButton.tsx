interface ButtonProps {
  text: string;
  onClick: () => void;
}

interface RoundedButtonProps extends ButtonProps {
  classNames?: string[];
  style?: object;
}

export const RoundedButton = (button: RoundedButtonProps) => {
  return (
    <input
      className={
        "p-2 px-8 drop-shadow-2xl rounded-full tracking-wider text-white font-bold  " +
        button.classNames?.join(" ")
      }
      style={button.style}
      type="button"
      onClick={button.onClick}
      value={button.text}
    />
  );
};

export const RoundedButtonPrimary = (button: ButtonProps) => {
  return (
    <RoundedButton
      text={button.text}
      onClick={button.onClick}
      classNames={["bg-green-500", "hover:bg-green-600"]}
    />
  );
};

export const RoundedButtonSecondary = (button: ButtonProps) => {
  return (
    <RoundedButton
      text={button.text}
      onClick={button.onClick}
      classNames={["bg-blue-500", "hover:bg-blue-600"]}
    />
  );
};

export default RoundedButton;
