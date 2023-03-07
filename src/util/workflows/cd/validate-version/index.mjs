import assert from 'node:assert';
import { createRegExp, exactly, digit, oneOrMore } from 'magic-regexp';

const VERSION_DEFINITION = createRegExp(
  exactly('v')
    .and(oneOrMore(digit).groupedAs('major'))
    .and(exactly('.').and(oneOrMore(digit).groupedAs('minor')))
    .and(exactly('.').and(oneOrMore(digit).groupedAs('patch')))
);

console.log('----- valid version syntax ---------------');
console.log('v0.1.0');
console.log('v0.1.0-SNAPSHOT');
console.log(VERSION_DEFINITION);
console.log('------------------------------------------');

// samples for valid versions
assert.equal(VERSION_DEFINITION.test(true, 'v0.1.0'));
assert.equal(VERSION_DEFINITION.test(true, 'v0.1.0-SNAPSHOT'));

// samples for invalid versions
assert.equal(VERSION_DEFINITION.test(false, '0.1'));
assert.equal(VERSION_DEFINITION.test(false, '0.1.0'));
assert.equal(VERSION_DEFINITION.test(false, '0-1-0'));
