import {useMutation} from "@tanstack/react-query";
import {axiosInstance} from "@/api/axiosInstance.ts";
import {RequestTopicRecords, ResponseTopicRecords} from "@/api/types.ts";
import {AxiosError} from "axios";


async function getTopicRecord(req: RequestTopicRecords) {
    const data = await axiosInstance.post<ResponseTopicRecords>("/topic/records", req);

    return data.data;
}

export const useGetTopicRecords = (req: RequestTopicRecords) => useMutation<ResponseTopicRecords, AxiosError, RequestTopicRecords>({
    queryKey: ['topics/records', JSON.stringify(req)],
    mutationFn: (req) => getTopicRecord(req),
})