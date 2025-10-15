export const getTagColor = (tag: string) => {
  const colors = [
    "from-blue-500 to-blue-600",
    "from-purple-500 to-purple-600",
    "from-pink-500 to-pink-600",
    "from-red-500 to-red-600",
    "from-orange-500 to-orange-600",
    "from-yellow-500 to-yellow-600",
    "from-green-500 to-green-600",
    "from-teal-500 to-teal-600",
    "from-cyan-500 to-cyan-600",
    "from-indigo-500 to-indigo-600",
  ];
  const index = tag.split("").reduce((acc, char) => acc + char.charCodeAt(0), 0);
  return colors[index % colors.length];
};
