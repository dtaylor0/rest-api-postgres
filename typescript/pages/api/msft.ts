// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import { db } from "@/db";
import { hist_msft } from "@/db/schema";
import type { NextApiRequest, NextApiResponse } from "next";

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse,
) {
  const queryRes = await db.select().from(hist_msft);
  res.status(200).json(queryRes);
}
