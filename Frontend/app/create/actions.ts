"use server";

import { useWriteContract } from "wagmi";
import { ABIS } from "../abis";

export async function createNFT(prevState, formData) {

  const { writeContractAsync, failureReason } = useWriteContract()

  const data = {
    title: formData.get("title"),
    summary: formData.get("summary"),
    content: formData.get("content"),
    topic: formData.get('topic'),
    grade: formData.get("grade"),
    subject: formData.get("subject"),
    price: formData.get('price')
  };
  const tokenId = await writeContractAsync({
    abi: ABIS,
    address: process.env.NEXT_PUBLIC_CONTRACT_ADDRESS as `0x${string}`,
    functionName: "createInnovation",
    args: [
        '1',
        '',
        '1',
        data.price,
        true
    ]
  })

  console.log(data, tokenId);
  return {};
}
