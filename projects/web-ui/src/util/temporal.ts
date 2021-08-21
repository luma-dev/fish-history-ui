import { Temporal } from '@js-temporal/polyfill';

export const firstDayOfMonth = (date: Temporal.PlainDate): Temporal.PlainDate => {
  return date.with({ day: 1 });
};

export const lastDayOfMonth = (date: Temporal.PlainDate): Temporal.PlainDate => {
  return date.with({ day: 1 }).add({ months: 1 }).subtract({ days: 1 });
};

export const firstDayOfWeek = (date: Temporal.PlainDate, dayStart = 0): Temporal.PlainDate => {
  return date.subtract({ days: (7 + date.dayOfWeek - dayStart) % 7 });
};

export const lastDayOfWeek = (date: Temporal.PlainDate, dayStart = 0): Temporal.PlainDate => {
  return date.add({ days: (6 + dayStart - date.dayOfWeek) % 7 });
};
