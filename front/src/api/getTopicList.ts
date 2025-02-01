import {useQuery} from "@tanstack/react-query";
import {axiosInstance} from "./axiosInstance.ts";
import {AxiosError} from "axios";

async function getTopics() {
    const data = await axiosInstance.get<string[]>("/topic/list");
    return data.data;
}

export const useTopicList = () => useQuery<string[], AxiosError>({queryKey: ['topics'], queryFn: getTopics})