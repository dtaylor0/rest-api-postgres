import { pgEnum, pgTable as table } from "drizzle-orm/pg-core";
import * as t from "drizzle-orm/pg-core";

export const rolesEnum = pgEnum("roles", ["app"]);

export const hist_msft = table(
  "hist_msft",
  {
    id: t.serial().primaryKey(),
    date: t.date().notNull(),
    close: t.numeric().notNull(),
    volume: t.text().notNull(),
    high: t.numeric().notNull(),
    low: t.numeric().notNull(),
  },
);
