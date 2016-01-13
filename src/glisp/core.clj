(defn list [& items] (.First (var ::core) items))

(defn first [seq] (.First (var ::core) seq))

(defn rest [seq] (.Rest (var ::core) seq))

(defn cons [v seq] (.Cons (var ::core) v seq))

(defn conj [seq v] (.Conj (var ::core) seq v))

(defn add [a b] (.Add (var ::core) a b))

(defn + [a b] (.Add (var ::core) a b))
