const axios = require('axios')
const { baseUrl } = require('../config')
const { normalGroup, normalGroup2 } = require('../__mock__/group.mock')
const { mockLoginUser, mockLoginUser2 } = require('../__mock__/user.mock')

const instance = axios.create({
  withCredentials: true,
  baseURL: baseUrl,
})

module.exports = {
  groupTests: () =>
    describe('group calls', () => {
      it('create group not authed', async () => {
        expect.assertions(1)
        try {
          await instance.post('group/create', normalGroup)
        } catch (err) {
          expect(err.response.status).toBe(401)
        }
      })

      describe('authed routes', () => {
        let user1, user2
        let groups, group1, group2
        beforeAll(async () => {
          user1 = await instance.post('user/login', mockLoginUser)
          user2 = await instance.post('user/login', mockLoginUser2)

          user1.cookie = user1.headers['set-cookie'][0].split(';')[0]
          user2.cookie = user2.headers['set-cookie'][0].split(';')[0]
        })

        it('create group authed', async () => {
          const res = await instance.post('group/create', normalGroup, {
            headers: { Cookie: user1.cookie },
            credentials: 'include',
          })
          expect(res.status).toBe(201)

          const res2 = await instance.post('group/create', normalGroup2, {
            headers: { Cookie: user2.cookie },
            credentials: 'include',
          })
          expect(res2.status).toBe(201)
        })

        it('check groups', async () => {
          const res = await instance.get('groups', {
            headers: { Cookie: user1.cookie },
          })
          expect(res.status).toBe(200)
          groups = res.data
          group1 = groups.filter((el) => el.title !== normalGroup.title)[0]
          group2 = groups.filter((el) => el.title !== normalGroup2.title)[0]

          const titles = groups.map((el) => el.title)
          expect(titles.includes(normalGroup.title)).toBe(true)
          expect(titles.includes(normalGroup2.title)).toBe(true)
        })

        it('group join', async () => {
          const res = await instance.post(
            'group/join',
            { groupID: group2.groupID },
            {
              headers: { Cookie: user1.cookie },
              credentials: 'include',
            }
          )
          expect(res.status).toBe(201)
        })

        it('group join admin', async () => {
          expect.assertions(1)
          try {
            await instance.post(
              'group/join',
              { groupID: group1.groupID },
              {
                headers: { Cookie: user1.cookie },
                credentials: 'include',
              }
            )
          } catch (err) {
            expect(err.response.status).toBe(400)
          }
        })

        it('group join another', async () => {
          const res = await instance.post(
            'group/join',
            { groupID: group2.groupID },
            {
              headers: { Cookie: user1.cookie },
              credentials: 'include',
            }
          )
          expect(res.status).toBe(201)
        })

        it('group check admin invites', async () => {
          const res = await instance.get('group/check-join', {
            headers: { Cookie: user2.cookie },
            credentials: 'include',
          })

          const res2 = await instance.get('group/check-join', {
            headers: { Cookie: user1.cookie },
            credentials: 'include',
          })

          console.log({ res })

          expect(res.status).toBe(200)

          console.log({ res2 })
        })
      })
    }),
}
