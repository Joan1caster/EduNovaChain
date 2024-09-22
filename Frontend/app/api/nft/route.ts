import { NextRequest } from "next/server";
import dayjs from "dayjs";
import { NFT, Topic } from "@/app/types";
import { cookies } from "next/headers";

const url = process.env.NEXT_PUBLIC_URL;
// const url = "https://f1f5-183-13-202-191.ngrok-free.app/api";

export async function GET(request: NextRequest) {
  const id = request.nextUrl.searchParams.get("id");
  const response = await fetch(`${url}/api/v1/nfts/${id}`, {
    method: "GET",
    cache: "no-store",
  });
  const data = await response.json();
  const nft = data.nft as NFT;

  const ipfsDataResponse = await fetch(
    `${request.nextUrl.origin}/api/ipfs?hash=${nft.MetadataURI}`
  );
  const ipfsData = await ipfsDataResponse.json();

  nft.Title = ipfsData.title;
  nft.Summary = ipfsData.summary;
  nft.Content = ipfsData.content;

  nft.CreatedAt = dayjs(nft.CreatedAt).format("YYYY-MM-DD");
  nft.UpdatedAt = dayjs(nft.UpdatedAt).format("YYYY-MM-DD");

  let CategoriesList: string[] = [];
  nft.Grades.map((grade: Topic) => {
    nft.Subjects.map((subject: Topic) => {
      CategoriesList.push(`${grade.Name}/${subject.Name}`);
    });
  });

  nft.Categories = CategoriesList.join(",");

  const historyResponse = await fetch(`${url}/api/v1/order/history`, {
    method: "POST",
    cache: "no-store",
    body: JSON.stringify({ nftId: nft.ID }),
  });

  const historyJsonData = await historyResponse.json();
  nft.HistoryList = historyJsonData.data;

  return Response.json(data.nft);
}

export async function POST(request: NextRequest) {
  const body = await request.json();
  const cookie = cookies();
  const Authorization = cookie.get("Authorization")?.value ?? "";

  request.headers.set("Authorization", Authorization);
  const response = await fetch(`${url}/api/v1/nfts/create`, {
    method: "POST",
    body: JSON.stringify(body),
    headers: {
      Authorization: Authorization,
    },
  });

  const data = await response.json();

  const orderResponse = await fetch(`${url}/api/v1/orders`, {
    method: "POST",
    body: JSON.stringify({
      nftId: data.id,
      price: body.price,
    }),
    headers: {
      Authorization: Authorization,
    },
  });

  const orderJson = await orderResponse.json();
  console.log(orderJson);

  return Response.json(data);
}

export async function PUT(request: NextRequest) {
  const body = await request.json();
  const response = await fetch(`${url}/api/v1/nfts/feature`, {
    method: "POST",
    body: JSON.stringify(body),
  });
  const data = await response.json();
  if (data.error) {
    return Response.json({
      code: 500,
      data,
    });
  } else {
    return Response.json({
      code: 200,
      data,
    });
  }
}
