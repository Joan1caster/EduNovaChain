"use client";

import { useState } from "react";
import { usePathname, useRouter } from "next/navigation";
import Image from "next/image";
import Link from "next/link";
import LoginButton from "./LoginButton";

export default function Header() {
  const pathname = usePathname();
  const router = useRouter();

  const [keywords, setKeywords] = useState("");

  return (
    <header className="px-10 py-4 mb-1">
      <div className="flex justify-between items-center text-primary">
        <Link href="/">
          <Image
            src="/images/slice/logo_01.png"
            alt="Vercel Logo"
            width={250}
            height={64}
            priority
          />
        </Link>
        <div className="flex-1 ml-6 flex gap-6 items-center">
          <Link
            href="/"
            className={`text-md ${pathname === "/" ? "text-black" : "text-primary-font_9"} hover:text-black`}
          >
            主页
          </Link>
          <Link
            href="/category"
            className={`text-md ${pathname === "/category" ? "text-black" : "text-primary-font_9"} hover:text-black`}
          >
            分类
          </Link>

          <div className="relative group/item text-md text-primary-font_9 cursor-pointer hover:text-black">
            发布
            <Image
              src={"/images/slice/AI.png"}
              alt={"AI.png"}
              width={20}
              height={42}
              className="absolute -top-1 left-9"
            />
            <div className="hidden static z-10 w-28 text-center group-hover/item:block group-hover/item:absolute bg-white shadow-lg rounded-md p-1">
              <Link
                className="block rounded-lg py-2 px-3 transition hover:bg-gray-100"
                href="/create"
              >
                <p className="text-primary-font_9 text-xs">发布创意</p>
              </Link>
              <Link
                className="block rounded-lg py-2 px-3 transition hover:bg-gray-100"
                href="/create"
              >
                <p className="text-primary-font_9 text-xs">AI生成创意</p>
              </Link>
            </div>
          </div>
        </div>
        <div className="w-1/3 bg-transparent flex justify-between">
          <div className="w-52 rounded bg-white py-1 pr-2 flex gap-2 justify-between">
            <input
              value={keywords}
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

        <div className="flex">
          <LoginButton email={""} />
        </div>
      </div>
    </header>
  );
}
