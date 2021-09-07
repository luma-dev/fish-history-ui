import type { Temporal } from '@js-temporal/polyfill';

export const firstDayOfMonth = (date: Temporal.PlainDate): Temporal.PlainDate => date.with({ day: 1 });

export const lastDayOfMonth = (date: Temporal.PlainDate): Temporal.PlainDate =>
  date.with({ day: 1 }).add({ months: 1 }).subtract({ days: 1 });

export const firstDayOfWeek = (date: Temporal.PlainDate, dayStart = 0): Temporal.PlainDate =>
  date.subtract({ days: (7 + date.dayOfWeek - dayStart) % 7 });

export const lastDayOfWeek = (date: Temporal.PlainDate, dayStart = 0): Temporal.PlainDate =>
  date.add({ days: (13 + dayStart - date.dayOfWeek) % 7 });
