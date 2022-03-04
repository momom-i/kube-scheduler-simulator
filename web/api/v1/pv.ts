import {
  V1PersistentVolume,
  V1PersistentVolumeList,
} from "@kubernetes/client-node";
import { k8sInstance } from "@/api/v1/index";

export const applyPersistentVolume = async (
  req: V1PersistentVolume,
  onError: (_: string) => void
) => {
  try {
    if (!req.metadata?.name) {
      onError("metadata.name is not provided");
      return;
    }
    req.kind = "PersistentVolume";
    req.apiVersion = "v1";
    const res = await k8sInstance.patch<V1PersistentVolume>(
      `/persistentvolumes/${req.metadata.name}?fieldManager=simulator`,
      req,
      { headers: { "Content-Type": "application/apply-patch+yaml" } }
    );
    return res.data;
  } catch (e: any) {
    onError("failed to applyPersistentVolume: " + e);
  }
};

export const listPersistentVolume = async (onError: (_: string) => void) => {
  try {
    const res = await k8sInstance.get<V1PersistentVolumeList>(
      `/persistentvolumes`,
      {}
    );
    return res.data;
  } catch (e: any) {
    onError("failed to listPersistentVolume: " + e);
  }
};

export const getPersistentVolume = async (
  name: string,
  onError: (_: string) => void
) => {
  try {
    const res = await k8sInstance.get<V1PersistentVolume>(
      `/persistentvolumes/${name}`,
      {}
    );
    return res.data;
  } catch (e: any) {
    onError("failed to getPersistentVolume: " + e);
  }
};

export const deletePersistentVolume = async (
  name: string,
  onError: (_: string) => void
) => {
  try {
    const res = await k8sInstance.delete(`/persistentvolumes/${name}`, {});
    return res.data;
  } catch (e: any) {
    onError("failed to deletePersistentVolume: " + e);
  }
};
