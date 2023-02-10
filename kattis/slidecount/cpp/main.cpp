#include <bits/stdc++.h>

void solve(const std::vector<int64_t> &elems, int64_t windowc) {
  int64_t s = 1, e = 1;
  std::vector<int64_t> freqs(elems.size());

  std::vector<int64_t> sums(elems.size());
  for (size_t i = 0; i < elems.size(); ++i) {
    sums[i] = (i == 0 ? 0 : sums[i - 1]) + elems[i];
  }

  auto sum = [sums](int64_t i, int64_t j) {
    if (i > j) {
      return (int64_t)0;
    }
    if (i == 0) {
      return sums[j - 1];
    }
    return sums[j - 1] - sums[i - 1];
  };

  auto slid = [&]() {
    if (s > e) {
      return;
    }

    for (size_t i = s - 1; i < e; ++i) {
      ++freqs[i];
    }
  };

  slid();
  while (s <= elems.size()) {
    if (e + 1 > elems.size() || sum(s - 1, e + 2 - 1) > windowc) {
      s++;
    } else {
      e++;
    }
    slid();
  }

  for (auto f : freqs) {
    std::cout << f << "\n";
  }
}

int main() {
  int64_t nelems, windowc;
  std::cin >> nelems >> windowc;
  std::vector<int64_t> elems(nelems);
  for (auto &w : elems) {
    std::cin >> w;
  }
  solve(elems, windowc);
  return 0;
}
