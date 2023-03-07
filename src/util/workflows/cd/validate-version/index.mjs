import assert from 'node:assert';
import { createRegExp, exactly, digit, oneOrMore } from 'magic-regexp';

const VERSION_REGEX = createRegExp(
  exactly('v')
    .and(oneOrMore(digit).groupedAs('major'))
    .and(exactly('.').and(oneOrMore(digit).groupedAs('minor')))
    .and(exactly('.').and(oneOrMore(digit).groupedAs('patch')))
);

console.log('----- given version ----------------------');
console.log(process.env.CHECK_VERSION);
console.log('valid = ' + VERSION_REGEX.test(process.env.CHECK_VERSION));
console.log('----- valid version syntax ---------------');
console.log('v0.1.0');
console.log('v0.1.0-SNAPSHOT');
console.log(VERSION_REGEX);
console.log('------------------------------------------');

// samples for valid versions
assert.equal(true, VERSION_REGEX.test('v0.1.0'));
assert.equal(true, VERSION_REGEX.test('v0.1.0-SNAPSHOT'));

// samples for invalid versions
assert.equal(false, VERSION_REGEX.test('0.1'));
assert.equal(false, VERSION_REGEX.test('0.1.0'));
assert.equal(false, VERSION_REGEX.test('0-1-0'));

// validate given version
assert.equal(true, VERSION_REGEX.test(process.env.CHECK_VERSION));
