import { useQuery } from "@tanstack/react-query";
import { getUnreadCount } from "@/lib/api/notification";

export const notificationKeys = {
  unreadCount: () => ["notifications", "unread-count"] as const,
};

export function useUnreadCount() {
  return useQuery({
    queryKey: notificationKeys.unreadCount(),
    queryFn: getUnreadCount,
  });
}
