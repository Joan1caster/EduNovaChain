import { getDefaultConfig } from "@rainbow-me/rainbowkit";
import { hardhat, polygonZkEvmTestnet, sepolia } from "wagmi/chains";

export const config = getDefaultConfig({
  appName: "RainbowKit App",
  projectId: "237dd1f5f9d2e2eae662f42e45220b2b",
  chains: [sepolia],
  ssr: true,
});
