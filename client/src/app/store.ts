import { configureStore, ThunkAction, Action } from "@reduxjs/toolkit";
import { rtkQueryApi } from "api/rtkQueryApi";
import { ClientDtoWithType } from "models/clients";
import { TypedUseSelectorHook, useDispatch, useSelector } from "react-redux";
// import counterReducer from "../features/List/counterSlice";

export const store = configureStore({
  reducer: {
    [rtkQueryApi.reducerPath]: rtkQueryApi.reducer,
    // counter: counterReducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({ serializableCheck: false }).concat([
      rtkQueryApi.middleware,
    ]),
});

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;
export type AppThunk<ReturnType = void> = ThunkAction<
  ReturnType,
  RootState,
  unknown,
  Action<string>
>;

export const useAppDispatch = () => useDispatch<AppDispatch>();
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector;

export const tstore: ClientDtoWithType[] = [];
