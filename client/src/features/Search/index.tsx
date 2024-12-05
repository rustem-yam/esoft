import React, { useEffect, useMemo, useState } from "react";
import {
  Backdrop,
  Grid,
  SpeedDial,
  SpeedDialAction,
  SpeedDialIcon,
  TextField,
  ToggleButton,
  ToggleButtonGroup,
} from "@mui/material";
import { useNavigate } from "react-router-dom";
import { useDebouncedCallback } from "use-debounce";
import { entities } from "app/entities";
import SearchItem from "./SearchItem";
import { useGetClientsQuery } from "api/common/clientsApi";

function Search() {
  const navigate = useNavigate();

  const { data: clientsResponse = { clients: [] } } = useGetClientsQuery();

  const allData = useMemo(() => {
    return {
      clients: clientsResponse.clients,
    };
  }, [clientsResponse]);

  const [isActionsOpened, setActionsOpened] = useState(false);
  const [filters, setFilters] = useState<string[]>([]);
  const [search, setSearch] = useState<string>("");

  const handleFilters = (
    event: React.MouseEvent<HTMLElement>,
    newFilters: string[]
  ) => {
    setFilters(newFilters);
  };

  const filteredCards = useMemo(() => {
    const combinedData = Object.entries(allData)
      .filter(([key]) => filters.includes(key) || !filters.length)
      .flatMap(([key, data]) =>
        data.map((item) => ({
          ...item,
          _type: key,
        }))
      );

    return combinedData.filter((item) =>
      Object.values(item).some((value) =>
        String(value).toLowerCase().includes(search.toLowerCase())
      )
    );
  }, [allData, filters, search]);

  const handleSearch = useDebouncedCallback((newSearch: string) => {
    setSearch(newSearch);
  }, 500);

  const toggleButtons = useMemo(
    () =>
      Object.keys(allData).map((key) => (
        <ToggleButton key={key} value={key}>
          {entities[key]?.label || key}
        </ToggleButton>
      )),
    [allData, entities]
  );

  const actionButtons = useMemo(
    () =>
      Object.keys(entities).map((key: string) => (
        <SpeedDialAction
          key={key}
          icon={entities[key].icon}
          tooltipTitle={entities[key].label}
          onClick={() => navigate("/form/" + key)}
        />
      )),
    [entities, navigate]
  );

  return (
    <>
      <TextField
        label="Search"
        id="search-input"
        fullWidth
        onChange={(e) => handleSearch(e.target.value)}
      />

      <ToggleButtonGroup
        value={filters}
        onChange={handleFilters}
        aria-label="filter entities"
        sx={{ my: 2 }}
      >
        {toggleButtons}
      </ToggleButtonGroup>

      <Grid container spacing={4}>
        {filteredCards.length === 0 ? (
          <div>No results found</div>
        ) : (
          filteredCards.map((card) => (
            <Grid item key={card.id} xs={12} sm={6} md={4}>
              <SearchItem data={card} type={card._type} />
            </Grid>
          ))
        )}
      </Grid>

      <Backdrop open={isActionsOpened} />
      <SpeedDial
        ariaLabel="Add"
        sx={{ position: "fixed", bottom: 16, right: 16 }}
        icon={<SpeedDialIcon />}
        onClose={() => setActionsOpened(false)}
        onOpen={() => setActionsOpened(true)}
        open={isActionsOpened}
      >
        {actionButtons}
      </SpeedDial>
    </>
  );
}

export default Search;
