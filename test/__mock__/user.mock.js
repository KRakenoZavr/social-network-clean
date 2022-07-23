const mockNormalUser1 = {
  email: 'fUser@mail.com',
  firstName: 'vasya',
  lastName: 'pupkin',
  password: 'lolkek',
  DateOfBirth: '2000-01-02T15:04:05Z',
}

const mockNormalUser2 = {
  email: 'sUser@mail.com',
  firstName: 'vasya',
  lastName: 'pupkin',
  password: 'lolkek2',
  DateOfBirth: '2000-01-02T15:04:05Z',
}

const mockNormalUserPrivate = {
  email: 'pUser@mail.com',
  firstName: 'vasya',
  lastName: 'pupkin',
  password: 'lolkek3',
  DateOfBirth: '2000-01-02T15:04:05Z',
  isPrivate: true,
}

const mockLoginUser = {
  email: 'fUser@mail.com',
  password: 'lolkek',
}

const mockBadLoginUser = {
  email: 'asdasd@mail.com',
  password: 'lol',
}

const mockBadLoginUser2 = {
  email: 'fUser@mail.com',
  password: 'lolke',
}

module.exports = {
  mockBadLoginUser,
  mockBadLoginUser2,
  mockLoginUser,
  mockNormalUser1,
  mockNormalUser2,
  mockNormalUserPrivate,
}
