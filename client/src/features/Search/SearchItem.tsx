import { Card, CardContent, CardHeader, Chip, Typography } from "@mui/material";
import { Link } from "react-router-dom";
import { entities } from "app/entities"; // Сущности
import { useMemo } from "react";

// Общий тип данных, который поддерживает любые сущности
type Props<T> = {
  data: T;
  type: string; // Тип сущности (например, "clients", "agents" и т.д.)
};

function SearchItem<T extends { id: number }>(props: Props<T>) {
  const { data, type } = props;

  // Извлекаем значения для заголовка на основе type и hFieldSubstr
  const valuesHeader = useMemo(() => {
    const hFieldSubstr = entities[type]?.hFieldSubstr || "";
    return Object.keys(data)
      .filter((key) => key.indexOf(hFieldSubstr) !== -1)
      .map((key) => data[key as keyof T]);
  }, [data, type]);

  // Извлекаем ключи для тела карточки, исключая id и заголовочные ключи
  const keysBody = useMemo(() => {
    const hFieldSubstr = entities[type]?.hFieldSubstr || "";
    return Object.keys(data).filter(
      (key) => key !== "id" && key.indexOf(hFieldSubstr) === -1
    );
  }, [data, type]);

  return (
    <Card
      component={Link}
      to={`/form/${type + "/" + data.id}`}
      sx={{
        height: "100%",
        display: "flex",
        flexDirection: "column",
        textDecoration: "none",
      }}
    >
      <CardHeader
        title={`${valuesHeader.join(" ")}`}
        action={
          <Chip
            color={entities[type]?.color || "default"}
            label={type}
            sx={{ cursor: "pointer" }}
          />
        }
      />
      <CardContent sx={{ flexGrow: 1 }}>
        {keysBody.map((key) => (
          <Typography key={data.id + key}>
            <b>{key}: </b>
            {String(data[key as keyof T])}
          </Typography>
        ))}
      </CardContent>
    </Card>
  );
}

export default SearchItem;
