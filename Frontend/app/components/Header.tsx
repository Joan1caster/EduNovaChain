"use client";

import { usePathname } from "next/navigation";
import Image from "next/image";
import Link from "next/link";
import LoginButton from "./LoginButton";

export default function Header() {
  const pathname = usePathname();

  return (
    <header className="px-2 py-4 mb-1 bg-gray-100">
      <div className="flex justify-between items-center">
        <Link href="/">
          {/* <Image
        src="/vercel.svg"
        alt="Vercel Logo"
        width={100}
        height={24}
        priority
      /> */}
          <img
            className="mx-auto h-5 w-auto inline-block mr-2"
            src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=600"
            alt="Your Company"
          />
          EduNovaChain
        </Link>
        <div className="flex-1 ml-4 flex gap-4 items-center">
          <Link
            href="/"
            className={`text-sm ${pathname === "/" ? "text-black" : "text-gray-500"} hover:text-black`}
          >
            主页
          </Link>
          <Link
            href="/category"
            className={`text-sm ${pathname === "/category" ? "text-black" : "text-gray-500"} hover:text-black`}
          >
            分类
          </Link>

          <div className="relative group/item text-sm text-gray-500 cursor-pointer hover:text-black">
            发布
            <div className="hidden static z-10 w-28 text-center group-hover/item:block group-hover/item:absolute bg-white shadow-lg rounded-md p-1">
              <Link
                className="block rounded-lg py-2 px-3 transition hover:bg-gray-100"
                href="/create"
              >
                <p className="text-gray-500 text-xs">发布创意</p>
              </Link>
              <Link
                className="block rounded-lg py-2 px-3 transition hover:bg-gray-100"
                href="/create"
              >
                <p className="text-gray-500 text-xs">AI生成创意</p>
              </Link>
            </div>
          </div>
        </div>
        <div className="w-1/2 bg-transparent flex justify-between">
          <div className="w-52 rounded bg-white py-1 px-4 flex gap-2 justify-between">
            <input
              placeholder="搜索"
              className="text-sm placeholder:text-slate-400 hover:outline-none focus:outline-none"
            />
            <Image
              src="/images/svg/search.svg"
              alt="search"
              width={16}
              height={16}
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
