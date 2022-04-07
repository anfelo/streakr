import { signIn, logout, onCurrentUserChanges, getDocument } from './firebase';

export async function signInWithEmailAndPassword(
  email: string,
  password: string
): Promise<{ status: number; body: any }> {
  try {
    const res = await signIn(email, password);
    const userData = await getDocument('users', res.body.uid);
    return userData;
  } catch (error) {
    return { status: 400, body: {} };
  }
}

export async function signOut(): Promise<{ status: number; body: any }> {
  try {
    await logout();
    return {
      status: 200,
      body: null,
    };
  } catch (error) {
    return { status: 400, body: {} };
  }
}

export async function getCurrentUserChanges(cb: Function) {
  onCurrentUserChanges(async (user) => {
    if (user && user.uid) {
      const userData = await getDocument('users', user.uid);
      return cb(userData);
    }
    cb({ status: 200, body: null });
  });
}
