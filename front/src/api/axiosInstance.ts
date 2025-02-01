import axios, {AxiosError, AxiosResponse, isAxiosError} from "axios";

const BASE_URL = "/api";


export const axiosInstance = axios.create({baseURL: BASE_URL});

// axiosInstance.interceptors.response.use(
//     async (response) => {
//         if (!response) return response;
//         const {
//             config,
//             data: { resCode, userMessage },
//             request,
//         } = response as AxiosResponse;
//
//         if (
//             resCode !== 200 &&
//             !config.headers['Skip-Notify'] &&
//             response.headers['content-type'] !== 'application/pdf' &&
//             response.headers['content-type'] !== 'application/octet-stream' &&
//             response.headers['content-type'] !==
//             'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet'
//         ) {
//             // TODO: добавить зависимость от ui-kit
//
//         }
//
//         if (
//             resCode !== 200 &&
//             response.headers['content-type'] !== 'application/octet-stream' &&
//             response.headers['content-type'] !== 'application/pdf' &&
//             response.headers['content-type'] !==
//             'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet'
//         ) {
//             throw new AxiosError(userMessage, 'error', config, request, response);
//         }
//
//         return response;
//     },
//     (error) => {
//         if (isAxiosError(error)) {
//
//         } else {
//             const axiosError = new AxiosError(
//                 error?.message,
//                 undefined,
//                 undefined,
//             );
//             axiosError.cause = error;
//             error = axiosError;
//         }
//
//         return Promise.reject(error);
//     },
// );