const mockNormalUser1 = {
  email: 'asdasd@mail.com',
  firstName: 'vasya',
  lastName: 'pupkin',
  password: 'lolkek',
  DateOfBirth: '2000-01-02T15:04:05Z',
}

const mockNormalUser2 = {
  email: 'asdasdasd@mail.com',
  firstName: 'vasya',
  lastName: 'pupkin',
  password: 'lolkek',
  DateOfBirth: '2000-01-02T15:04:05Z',
}

const mockLoginUser = {
  email: 'asdasd@mail.com',
  password: 'lolkek',
}

const mockBadLoginUser = {
  email: 'asdasd@mail.com',
  password: 'lol',
}

const mockBadLoginUser2 = {
  email: 'asd@mail.com',
  password: 'lolkek',
}

module.exports = {
  mockBadLoginUser,
  mockBadLoginUser2,
  mockLoginUser,
  mockNormalUser1,
  mockNormalUser2,
}
