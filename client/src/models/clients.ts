type ClientDto = {
  id: number;
  name: string;
  surname: string;
  patronymic: string;
  email?: string;
  phone?: string;
};

type ClientDtoWithType = ClientDto & { type: string };

type AddClientMutationRequestDto = Omit<ClientDto, "id">;

export type { ClientDto, ClientDtoWithType, AddClientMutationRequestDto };
