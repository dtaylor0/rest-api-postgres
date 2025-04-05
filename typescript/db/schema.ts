import { pgEnum, pgTable as table } from "drizzle-orm/pg-core";
import * as t from "drizzle-orm/pg-core";

export const rolesEnum = pgEnum("roles", ["app"]);

export const hist_msft = table("hist_msft", {
  id: t.serial().primaryKey(),
  date: t.date({ mode: "date" }).notNull(),
  close: t.numeric({ mode: "number" }).notNull(),
  open: t.numeric({ mode: "number" }).notNull(),
  volume: t.bigint({ mode: "number" }).notNull(),
  high: t.numeric({ mode: "number" }).notNull(),
  low: t.numeric({ mode: "number" }).notNull(),
});
