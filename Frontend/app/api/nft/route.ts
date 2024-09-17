import { NextRequest } from "next/server";
import dayjs from "dayjs";
import { NFT, Topic } from "@/app/types";

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
  const nft = data.nft as NFT

  nft.CreatedAt = dayjs(nft.CreatedAt).format("YYYY-MM-DD");
  nft.UpdatedAt = dayjs(nft.UpdatedAt).format("YYYY-MM-DD");

  let CategoriesList: string[] = []
  nft.Grades.map((grade: Topic) => {
    nft.Subjects.map((subject: Topic) => {
      CategoriesList.push(`${grade.Name}/${subject.Name}`)
    })
  })

  nft.Categories = CategoriesList.join(',')
  
  return Response.json(data.nft);
}
