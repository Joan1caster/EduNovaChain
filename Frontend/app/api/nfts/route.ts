import { NFT } from "@/app/types";
import dayjs from "dayjs";
import { NextRequest } from "next/server";
const url = process.env.NEXT_PUBLIC_URL;
// const url = "https://f1f5-183-13-202-191.ngrok-free.app/api";

export async function GET(request: NextRequest) {
  const type = request.nextUrl.searchParams.get("type");
  const count = request.nextUrl.searchParams.get("count") ?? 10;
  const response = await fetch(`${url}/api/v1/nfts/${type}/${count}`, {
    method: "GET",
    cache: "no-store",
  });
  const data = await response.json();
  const idsFetch: any[] = [];
  let list: NFT[] = [];
  if (data.data.length > 0) {
    data.data.forEach((item: any) => {
      idsFetch.push(
        fetch(`${request.nextUrl.origin}/api/ipfs?hash=${item.MetadataURI}`)
      );
      item.CreatedAt = dayjs(item.CreatedAt).format("YYYY-MM-DD");
      item.UpdatedAt = dayjs(item.UpdatedAt).format("YYYY-MM-DD");
      item.ContentFeature = "rewrite";
      item.SummaryFeature = "rewrite";
    });
  }

  const promiseList = await Promise.all(idsFetch);
  const responseJsonList = await Promise.all(
    promiseList.map((item) => item.json())
  );
  responseJsonList.map((rep, i) => {
    list.push({
      ...data.data[i],
      Title: rep.title,
      Summary: rep.summary,
      Content: rep.content,
    });
  });

  return Response.json({
    count: data.count ?? 0,
    data: list,
  });
}
