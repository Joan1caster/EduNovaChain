"use client";

import { useAsyncEffect } from "ahooks";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { useEffect, useState } from "react";
import { useAccount } from "wagmi";
import Image from "next/image";
import { getSiweMessage } from "../services";
import { ConnectButton } from "@rainbow-me/rainbowkit";
import { Dialog, DialogPanel } from "@headlessui/react";

export default function LoginButton({ email }: { email: string }) {
  const [isOpen, setIsOpen] = useState(true);
  const { address, isConnected } = useAccount();
  useAsyncEffect(async () => {
    if (address && isConnected) {
      await getSiweMessage(address);
      console.log(address, isConnected);
    }
  }, [address, isConnected]);

  return (
    <>
      {isConnected ? (
        <ConnectButton accountStatus="address" />
      ) : (
        <>
          <button className="rounded-md bg-primary px-6 py-2 mx-auto font-semibold text-white shadow-sm hover:bg-primary focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary-light/50 disabled:bg-primary-light/50 disabled:cursor-not-allowed">
            Web3钱包
          </button>
          <Dialog
            open={isOpen}
            onClose={() => setIsOpen(false)}
            className="relative z-5"
          >
            <div className="fixed inset-0 flex w-screen h-screen items-center justify-center p-4 bg-black/40">
              <DialogPanel className="w-[880px] h-[432px] flex items-center space-y-4 bg-white rounded-lg overflow-hidden">
                <div className="w-[408px] h-full flex justify-center items-center bg-[url('/images/slice/login_bg.png')] bg-no-repeat bg-cover">
                  <Image
                    src={"/images/slice/logo_02.png"}
                    width={180}
                    height={96}
                    alt="logo"
                  />
                </div>
                <div className="px-10">
                  <p className="text-[40px]">Web3 钱包</p>
                  <p className="text-[#999] mt-6 mb-16">
                    最安全的自托管钱包，交易更快、更放心
                  </p>

                  <button className="rounded-md bg-primary px-6 py-2 mx-auto font-semibold text-white shadow-sm hover:bg-primary focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary-light/50 disabled:bg-primary-light/50 disabled:cursor-not-allowed">
                    <ConnectButton label="链接钱包" />
                  </button>
                </div>
              </DialogPanel>
            </div>
          </Dialog>
        </>
      )}
    </>
  );
}
