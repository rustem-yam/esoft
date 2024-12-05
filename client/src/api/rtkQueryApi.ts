import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

const rtkQueryApi = createApi({
  reducerPath: "api",
  baseQuery: fetchBaseQuery({
    baseUrl: `http://localhost:4000/api`,
    prepareHeaders: (headers) => {
      headers.set("Accept", "application/json");

      return headers;
    },
  }),
  endpoints: () => ({}),
});

export { rtkQueryApi };
