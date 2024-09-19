import { NextRequest } from "next/server";
import dayjs from "dayjs";
import { NFT, Topic } from "@/app/types";
import { cookies } from "next/headers";

const url = process.env.NEXT_PUBLIC_URL;
// const url = "https://f1f5-183-13-202-191.ngrok-free.app/api";

export async function GET(request: NextRequest) {
  const id = request.nextUrl.searchParams.get("id");
  console.log(id);
  const response = await fetch(`${url}/api/v1/nfts/${id}`, {
    method: "GET",
    cache: "no-store",
  });
  const data = await response.json();
  const nft = data.nft as NFT;

  nft.CreatedAt = dayjs(nft.CreatedAt).format("YYYY-MM-DD");
  nft.UpdatedAt = dayjs(nft.UpdatedAt).format("YYYY-MM-DD");

  let CategoriesList: string[] = [];
  nft.Grades.map((grade: Topic) => {
    nft.Subjects.map((subject: Topic) => {
      CategoriesList.push(`${grade.Name}/${subject.Name}`);
    });
  });

  nft.Categories = CategoriesList.join(",");

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

  return Response.json(data);
}

export async function PUT(request: NextRequest) {
  const body = await request.json();
  const response = await fetch(`${url}/api/v1/nfts/feature`, {
    method: "POST",
    body: JSON.stringify(body),
  });
  const data = await response.json();

  return Response.json({
    code: 200,
    data,
  });
}
