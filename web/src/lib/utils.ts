export const getAvatarNumber = (tag: string): number => {
  const index = tag
    .split("")
    .reduce((acc, char) => acc + char.charCodeAt(0), 0);
  return (index % 5) + 1;
};
