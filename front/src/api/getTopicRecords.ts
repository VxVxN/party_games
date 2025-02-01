import { useMutation } from "@tanstack/react-query";
import { axiosInstance } from "@/api/axiosInstance.ts";
import { RequestTopicRecords, ResponseTopicRecords } from "@/api/types.ts";
import { AxiosError } from "axios";

async function getTopicRecord(req: RequestTopicRecords) {
  const data = await axiosInstance.post<ResponseTopicRecords>(
    "/topic/records",
    { ...req, page_size: req.page_size || 30 },
  );

  return data.data;
}

export const useGetTopicRecords = () =>
  useMutation<ResponseTopicRecords, AxiosError, RequestTopicRecords>({
    mutationFn: (req) => getTopicRecord(req),
  });
