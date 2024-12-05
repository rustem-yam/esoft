import { useForm } from "react-hook-form";
import { useAddClientMutation } from "api/common/clientsApi";
import { AddClientMutationRequestDto } from "models/clients";
import { Button } from "@mui/material";
import { TextFieldElement } from "react-hook-form-mui";

function ClientForm() {
  const {
    control,
    handleSubmit,
    formState: { errors },
  } = useForm<AddClientMutationRequestDto>();
  const [addClient] = useAddClientMutation();

  const onSubmit = (data: AddClientMutationRequestDto) => {
    console.log("Form data:", data);
    addClient(data);
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <TextFieldElement
        fullWidth
        sx={{ my: 1 }}
        name="name"
        label="Last Name"
        required
        control={control}
        error={!!errors.name}
        helperText={errors.name ? "Last name is required" : ""}
      />
      <TextFieldElement
        fullWidth
        sx={{ my: 1 }}
        name="surname"
        label="First Name"
        required
        control={control}
        error={!!errors.surname}
        helperText={errors.surname ? "First name is required" : ""}
      />
      <TextFieldElement
        fullWidth
        sx={{ my: 1 }}
        name="patronymic"
        label="Middle Name"
        control={control}
        error={!!errors.patronymic}
        helperText={errors.patronymic ? "Middle name is required" : ""}
      />
      <TextFieldElement
        fullWidth
        sx={{ my: 1 }}
        name="email"
        label="Email"
        required
        type="email"
        control={control}
        error={!!errors.email}
        helperText={errors.email ? "Valid email is required" : ""}
      />
      <TextFieldElement
        fullWidth
        sx={{ my: 1 }}
        name="telephone"
        label="Phone"
        required
        type="tel"
        control={control}
        error={!!errors.telephone}
        helperText={errors.telephone ? "Phone number is required" : ""}
      />

      <Button variant="contained" fullWidth sx={{ my: 2 }} type="submit">
        Submit
      </Button>
    </form>
  );
}

export default ClientForm;
