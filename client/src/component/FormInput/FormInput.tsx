import { toTitleCase } from "../../common/functions";

export interface FormInputProps {
  type: "button" | "text" | "password" | "datetime" | "email";
  name: string;
  label: string;
  id: string;
  value?: string;
  onChange?: (e: React.ChangeEvent<HTMLInputElement>) => void;
}

const FormInput = (input: FormInputProps) => {
  return (
    <div className="flex-col  m-2">
      <label className="block bg-red" htmlFor={input.id}>
        {toTitleCase(input.label)}
      </label>
      <input
        value={input.value}
        onChange={input.onChange}
        className="p-2 px-4 border w-full border-gray-300 rounded-xl"
        type={input.type}
        id={input.id}
        name={input.name}
      ></input>
    </div>
  );
};

export default FormInput;
