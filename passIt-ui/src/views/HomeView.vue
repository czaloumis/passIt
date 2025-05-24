<script setup lang="ts">
import Hero from '@/components/Hero.vue';
import HomeCards from '@/components/HomeCards.vue';
import JobListing from '@/components/JobListing.vue';
</script>

<template>
    <Hero title="Test Vue" subtitle="Test subtitle" />
    <HomeCards />
    <JobListing :limit="3" :showAll="true" />
</template>