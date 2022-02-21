import { V1PriorityClass, V1PriorityClassList } from "@kubernetes/client-node";
import { k8sSchedulingInstance } from "@/api/v1/index";

export const applyPriorityClass = async (
  req: V1PriorityClass,
  onError: (_: string) => void
) => {
  try {
    if (!req.metadata?.name) {
      onError("metadata.name is not provided");
      return;
    }
    const res = await k8sSchedulingInstance.patch<V1PriorityClass>(
      `/priorityclasses/${req.metadata.name}?fieldManager=simulator`,
      req,
      { headers: { "Content-Type": "application/strategic-merge-patch+json" } }
    );
    return res.data;
  } catch (e: any) {
    try {
      const res = await createPriorityClass(req, onError);
      return res;
    } catch (e: any) {
      onError(e);
    }
  }
};

export const listPriorityClass = async () => {
  const res = await k8sSchedulingInstance.get<V1PriorityClassList>(
    `/priorityclasses`,
    {}
  );
  return res.data;
};

export const getPriorityClass = async (name: string) => {
  const res = await k8sSchedulingInstance.get<V1PriorityClass>(
    `/priorityclasses/${name}`,
    {}
  );
  return res.data;
};

export const deletePriorityClass = async (name: string) => {
  const res = await k8sSchedulingInstance.delete(
    `/priorityclasses/${name}`,
    {}
  );
  return res.data;
};

const createPriorityClass = async (
  req: V1PriorityClass,
  onError: (_: string) => void
) => {
  try {
    const res = await k8sSchedulingInstance.post<V1PriorityClass>(
      `/priorityclasses?fieldManager=simulator`,
      req
    );
    return res.data;
  } catch (e: any) {
    onError(e);
  }
};
