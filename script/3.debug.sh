#/bin/bash
NAME=$(kubectl get pod -n keti-system | grep -E 'exascale-web' | awk '{print $1}')

#echo "Exec Into '"$NAME"'"

#kubectl exec -it $NAME -n $NS /bin/sh
kubectl logs -f $NAME -n keti-system