import { Temporal } from '@js-temporal/polyfill';
import { firstDayOfMonth, firstDayOfWeek, lastDayOfMonth, lastDayOfWeek } from '../temporal';

describe('firstDayOfMonth', () => {
  it('returns first day of month', () => {
    expect(firstDayOfMonth(new Temporal.PlainDate(2000, 1, 1)).toString()).toBe(
      new Temporal.PlainDate(2000, 1, 1).toString(),
    );
    expect(firstDayOfMonth(new Temporal.PlainDate(2000, 1, 2)).toString()).toBe(
      new Temporal.PlainDate(2000, 1, 1).toString(),
    );
    expect(firstDayOfMonth(new Temporal.PlainDate(2000, 1, 30)).toString()).toBe(
      new Temporal.PlainDate(2000, 1, 1).toString(),
    );
    expect(firstDayOfMonth(new Temporal.PlainDate(2000, 1, 31)).toString()).toBe(
      new Temporal.PlainDate(2000, 1, 1).toString(),
    );
    expect(firstDayOfMonth(new Temporal.PlainDate(2000, 2, 1)).toString()).toBe(
      new Temporal.PlainDate(2000, 2, 1).toString(),
    );
    expect(firstDayOfMonth(new Temporal.PlainDate(2000, 2, 29)).toString()).toBe(
      new Temporal.PlainDate(2000, 2, 1).toString(),
    );
    expect(firstDayOfMonth(new Temporal.PlainDate(2000, 3, 1)).toString()).toBe(
      new Temporal.PlainDate(2000, 3, 1).toString(),
    );
    expect(firstDayOfMonth(new Temporal.PlainDate(2000, 12, 1)).toString()).toBe(
      new Temporal.PlainDate(2000, 12, 1).toString(),
    );
    expect(firstDayOfMonth(new Temporal.PlainDate(2000, 12, 31)).toString()).toBe(
      new Temporal.PlainDate(2000, 12, 1).toString(),
    );
  });
});

describe('lastDayOfMonth', () => {
  it('returns last day of month', () => {
    expect(lastDayOfMonth(new Temporal.PlainDate(2000, 1, 1)).toString()).toBe(
      new Temporal.PlainDate(2000, 1, 31).toString(),
    );
    expect(lastDayOfMonth(new Temporal.PlainDate(2000, 1, 31)).toString()).toBe(
      new Temporal.PlainDate(2000, 1, 31).toString(),
    );
    expect(lastDayOfMonth(new Temporal.PlainDate(2000, 2, 1)).toString()).toBe(
      new Temporal.PlainDate(2000, 2, 29).toString(),
    );
    expect(lastDayOfMonth(new Temporal.PlainDate(2000, 2, 29)).toString()).toBe(
      new Temporal.PlainDate(2000, 2, 29).toString(),
    );
    expect(lastDayOfMonth(new Temporal.PlainDate(2000, 12, 1)).toString()).toBe(
      new Temporal.PlainDate(2000, 12, 31).toString(),
    );
    expect(lastDayOfMonth(new Temporal.PlainDate(2000, 12, 31)).toString()).toBe(
      new Temporal.PlainDate(2000, 12, 31).toString(),
    );
  });
});

describe('firstDayOfWeek', () => {
  it('returns first day of week', () => {
    // It is Saturday(6) on 2020/1/1
    expect(firstDayOfWeek(new Temporal.PlainDate(2000, 1, 1)).toString()).toBe(
      new Temporal.PlainDate(1999, 12, 26).toString(),
    );
    expect(firstDayOfWeek(new Temporal.PlainDate(2000, 1, 1), 1).toString()).toBe(
      new Temporal.PlainDate(1999, 12, 27).toString(),
    );
    expect(firstDayOfWeek(new Temporal.PlainDate(2000, 1, 1), 6).toString()).toBe(
      new Temporal.PlainDate(2000, 1, 1).toString(),
    );
  });
});

describe('lastDayOfWeek', () => {
  it('returns last day of week', () => {
    expect(lastDayOfWeek(new Temporal.PlainDate(2000, 1, 1)).toString()).toBe(
      new Temporal.PlainDate(2000, 1, 1).toString(),
    );
    expect(lastDayOfWeek(new Temporal.PlainDate(2000, 1, 1), 1).toString()).toBe(
      new Temporal.PlainDate(2000, 1, 2).toString(),
    );
    expect(lastDayOfWeek(new Temporal.PlainDate(2000, 1, 1), 6).toString()).toBe(
      new Temporal.PlainDate(2000, 1, 7).toString(),
    );
    expect(lastDayOfWeek(new Temporal.PlainDate(2021, 2, 28)).toString()).toBe(
      new Temporal.PlainDate(2021, 3, 6).toString(),
    );
  });
});
