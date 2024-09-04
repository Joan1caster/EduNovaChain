"use client";
import { useState } from "react";
import { TagType } from "../types";

export default function Tab({
  data,
  onChange,
}: {
  data: TagType[];
  onChange: (item: TagType) => void;
}) {
  const [currentType, setCurrentType] = useState<TagType>(data[0]);

  const onSwitchType = (item: TagType) => {
    setCurrentType(item);
  };
  return (
    <ul className="flex w-auto justify-around rounded-md bg-primary-border/50 text-[0.75rem] *:text-[0.75rem] *:hover:cursor-pointer">
      {data.map((item) => (
        <li
          className={`px-3 py-1 rounded-md text-center ${currentType.key === item.key ? " bg-primary-border " : ""}`}
          onClick={() => onSwitchType(item)}
        >
          {item.name}
        </li>
      ))}
    </ul>
  );
}
