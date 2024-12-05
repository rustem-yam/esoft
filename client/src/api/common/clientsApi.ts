import { rtkQueryApi } from "api/rtkQueryApi";
import { AddClientMutationRequestDto, ClientDto } from "models/clients";

const clientsApi = rtkQueryApi.injectEndpoints({
  endpoints: (build) => ({
    getClients: build.query<ClientDto[], void>({
      query: () => ({
        url: "/clients",
        method: "GET",
      }),
    }),
    addClient: build.mutation<ClientDto, AddClientMutationRequestDto>({
      query: (body) => ({
        url: "/clients",
        method: "POST",
        body,
      }),
    }),
  }),
});

export const { useGetClientsQuery, useAddClientMutation } = clientsApi;
