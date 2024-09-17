"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import Image from "next/image";

export default function Search() {
  const router = useRouter();

  const [keywords, setKeywords] = useState("");
  return (
    <div className="w-1/3 bg-transparent flex justify-between">
      <div className="w-52 rounded bg-white py-1 pr-2 flex gap-2 justify-between">
        <input
          value={keywords}
          readOnly
          placeholder="搜索"
          className="text-xs placeholder:text-primary-font_C border-none outline-none shadow-none focus:shadow-none focus:ring-offset-0 focus:ring-0"
        />
        {/* <input
              type="text"
              placeholder="搜索"
              className="block w-full border-none text-gray-900 placeholder:text-gray-400 sm:text-xs focus:outline-none focus:shadow-none focus:border-none"
            ></input> */}
        <Image
          src="/images/svg/search.svg"
          alt="search"
          width={16}
          height={16}
          onClick={() => router.push(`/search?key=${keywords}`)}
        />
      </div>
    </div>
  );
}
