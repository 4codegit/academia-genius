export const useFilters = () => {
  const selectedTopics = ref<string[]>([])
  const topicsList = ['Механика','МКТ','Термодинамика','Электростатика','Магнетизм','Оптика','СТО','Квантовая'] as const
  const toggleTopic = (t: string) => {
    const i = selectedTopics.value.indexOf(t)
    if (i >= 0) {
      selectedTopics.value.splice(i, 1)
    } else {
      selectedTopics.value.push(t)
    }
  }
  const clearTopics = () => { selectedTopics.value = [] }
  return { selectedTopics, topicsList: readonly(topicsList), toggleTopic, clearTopics }
}
