import { initializeApp } from 'firebase/app';
import {
  getFirestore,
  getDoc,
  getDocs,
  setDoc,
  updateDoc,
  deleteDoc,
  doc,
  collection,
  orderBy,
  query,
} from 'firebase/firestore';
import {
  getAuth,
  signInWithEmailAndPassword,
  signOut,
  onAuthStateChanged,
} from 'firebase/auth';
import { getStorage, ref, getDownloadURL, uploadBytes } from 'firebase/storage';

const firebaseConfig = {
  apiKey: 'AIzaSyAj0RcsICXur0I_8_L-F3uyRVLtOya4O-0',
  authDomain: 'streakr-df522.firebaseapp.com',
  projectId: 'streakr-df522',
  storageBucket: 'streakr-df522.appspot.com',
  messagingSenderId: '282419695019',
  appId: '1:282419695019:web:9bba9e087634887732d6ee',
};

const app = initializeApp(firebaseConfig);
export const db = getFirestore(app);
export const auth = getAuth(app);
export const storage = getStorage(app);

export async function getCollection(
  path: string,
  sortBy?: { key: string; direction: 'desc' | 'asc' }
): Promise<{ status: number; body: any }> {
  if (!path) return { status: 400, body: {} };
  const colRef = collection(db, path);
  let queryRef = query(colRef);
  if (sortBy) {
    queryRef = query(colRef, orderBy(sortBy.key, sortBy.direction));
  }
  const querySnapshot = await getDocs(queryRef);
  const colData: any[] = [];

  querySnapshot.forEach((doc) => {
    colData.push({
      id: doc.id,
      ...doc.data(),
    });
  });

  return {
    status: 200,
    body: [...colData],
  };
}

export async function getDocument(
  path: string,
  id: string
): Promise<{ status: number; body: any }> {
  if (!path) return { status: 400, body: {} };
  const docRef = doc(db, path, id);
  const docSnapshot = await getDoc(docRef);
  const data = docSnapshot.data();

  if (data) {
    return {
      status: 200,
      body: {
        id: docSnapshot.id,
        ...data,
      },
    };
  } else {
    return { status: 400, body: {} };
  }
}

export async function addOrUpdateDocument(
  path: string,
  data: any
): Promise<{ status: number; body: any }> {
  let docData;
  if (!data.id) {
    docData = {
      ...data,
      updated: new Date().toISOString(),
      created: new Date().toISOString(),
    };
    await setDoc(doc(db, path, data.slug), docData);
  } else {
    docData = {
      ...data,
      updated: new Date().toISOString(),
    };
    delete docData.id;
    await updateDoc(doc(db, path, data.id), docData);
  }
  return {
    status: 200,
    body: docData,
  };
}

export async function deleteDocument(path: string, id: string) {
  await deleteDoc(doc(db, path, id));
}

export async function uploadObject(path: string, file: File) {
  const storageRef = ref(storage, path);
  const snapshot = await uploadBytes(storageRef, file);
  return snapshot.metadata.name;
}

export async function getObjectUrl(path: string) {
  const storageRef = ref(storage, path);
  const contentUrl = await getDownloadURL(storageRef);
  return contentUrl;
}

export async function signIn(
  email: string,
  password: string
): Promise<{ status: number; body: any }> {
  const userCredential = await signInWithEmailAndPassword(
    auth,
    email,
    password
  );
  return {
    status: 200,
    body: userCredential.user,
  };
}

export async function logout() {
  return await signOut(auth);
}

export function onCurrentUserChanges(cb: Function) {
  onAuthStateChanged(auth, (user) => {
    cb(user);
  });
}
