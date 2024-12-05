type ClientDto = {
  id: number;
  name: string;
  surname: string;
  patronymic: string;
  email?: string;
  telephone?: string;
};

type GetClientsQueryRequestDto = {
  clients: ClientDto[];
};

type AddClientMutationRequestDto = Omit<ClientDto, "id">;

export type {
  ClientDto,
  GetClientsQueryRequestDto,
  AddClientMutationRequestDto,
};
