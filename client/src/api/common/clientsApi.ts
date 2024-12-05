import { rtkQueryApi } from "api/rtkQueryApi";
import {
  AddClientMutationRequestDto,
  ClientDto,
  GetClientsQueryRequestDto,
} from "models/clients";

const clientsApi = rtkQueryApi.injectEndpoints({
  endpoints: (build) => ({
    getClients: build.query<GetClientsQueryRequestDto, void>({
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
